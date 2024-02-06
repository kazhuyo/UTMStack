package com.park.utmstack.service.incident_response;

import com.park.utmstack.config.Constants;
import com.park.utmstack.domain.incident_response.UtmIncidentVariable;
import com.park.utmstack.repository.incident_response.UtmIncidentVariableRepository;
import com.park.utmstack.util.CipherUtil;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.util.CollectionUtils;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * Service Implementation for managing UtmIncidentVariable.
 */
@Service
@Transactional
public class UtmIncidentVariableService {

    private final Logger log = LoggerFactory.getLogger(UtmIncidentVariableService.class);

    private final UtmIncidentVariableRepository utmIncidentVariableRepository;

    public UtmIncidentVariableService(UtmIncidentVariableRepository utmIncidentVariableRepository) {

        this.utmIncidentVariableRepository = utmIncidentVariableRepository;
    }

    /**
     * Save a utmIncidentVariable.
     *
     * @param utmIncidentVariable the entity to save
     * @return the persisted entity
     */
    public UtmIncidentVariable save(UtmIncidentVariable utmIncidentVariable) {
        log.debug("Request to save UtmIncidentVariable : {}", utmIncidentVariable);
        if (utmIncidentVariable.isSecret()) {
            String currentValue = utmIncidentVariable.getVariableValue();
            utmIncidentVariable.setVariableValue(CipherUtil.encrypt(currentValue, System.getenv(Constants.ENV_ENCRYPTION_KEY)));
        }
        return utmIncidentVariableRepository.save(utmIncidentVariable);
    }

    /**
     * Get all the utmIncidentVariables.
     *
     * @param pageable the pagination information
     * @return the list of entities
     */
    @Transactional(readOnly = true)
    public Page<UtmIncidentVariable> findAll(Pageable pageable) {
        log.debug("Request to get all UtmIncidentVariables");
        return utmIncidentVariableRepository.findAll(pageable);
    }


    /**
     * Get one utmIncidentVariable by id.
     *
     * @param id the id of the entity
     * @return the entity
     */
    @Transactional(readOnly = true)
    public Optional<UtmIncidentVariable> findOne(Long id) {
        log.debug("Request to get UtmIncidentVariable : {}", id);
        return utmIncidentVariableRepository.findById(id);
    }

    /**
     * Get one utmIncidentVariable by name.
     *
     * @param variableName the name of the variable
     * @return the entity
     */
    @Transactional(readOnly = true)
    public Optional<UtmIncidentVariable> findByName(String variableName) {
        return utmIncidentVariableRepository.findByVariableName(variableName);
    }

    /**
     * Get all utmIncidentVariable by name in list.
     *
     * @param variableNames the names of the variables to list
     * @return the entity
     */
    @Transactional(readOnly = true)
    public List<UtmIncidentVariable> findByVariablesByNames(List<String> variableNames) {
        return utmIncidentVariableRepository.findAllByVariableNameIn(variableNames);
    }


    /**
     * Delete the utmIncidentVariable by id.
     *
     * @param id the id of the entity
     */
    public void delete(Long id) {
        log.debug("Request to delete UtmIncidentVariable : {}", id);
        utmIncidentVariableRepository.deleteById(id);
    }

    public String replaceVariablesInCommand(String command) {
        List<String> variableNames = new ArrayList<>();
        Pattern pattern = Pattern.compile("\\$\\[variables\\.(.*?)]");
        Matcher matcher = pattern.matcher(command);

        while (matcher.find()) {
            variableNames.add(matcher.group(1));
        }

        List<UtmIncidentVariable> variables = findByVariablesByNames(variableNames);
        if (CollectionUtils.isEmpty(variables))
            return command;

        for (UtmIncidentVariable var : variables) {
            String currentValue = var.getVariableValue();
            if (var.isSecret())
                currentValue = "$[" + var.getVariableName() + ":" + currentValue + "]";
            command = command.replace("$[variables." + var.getVariableName() + "]", currentValue);
        }

        return command;
    }
}
