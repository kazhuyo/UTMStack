from string import Template

from mutate.cloudIntegrations.integration import Integration
from mutate.util.module_enum import AZURE
from mutate.util.utils import get_module_group


def get_event_hubs(conf):
    if "EntityPath" in str(conf["eventHubConnection"]):
        name = conf["eventHubConnection"].split("EntityPath=")[1].split(";")[0]
        event_hubs = "{\"" + name + "\" => {\n"
        event_hubs += ("         event_hub_connection => \"{}\"\n"
                       "         storage_connection => \"{}\" \n"
                       "         consumer_group => \"{}\"\n"
                       "         storage_container => \"{}\"\n"
                       "").format(conf["eventHubConnection"],
                                  conf["storageConnection"],
                                  conf["consumerGroup"],
                                  conf["storageContainer"])
        return event_hubs + "       }\n     }"
    else:
        return None


class AzureIntegration(Integration):
    def __init__(self):
        Integration.__init__(self)

    def get_integration_config(self) -> str:
        """Implement interface, build string for
        integration"""
        azure = ""
        module = AZURE
        groups = get_module_group("AZURE")
        if groups is not None and len(groups) > 0:
            for group in groups:
                azure_configs = self.get_input_integration(module, 'eventHubConnection', group)
                if azure_configs is not None:
                    config = get_event_hubs(azure_configs)
                    if config is not None:
                        azure += "azure_event_hubs {\n" \
                                 "   config_mode => \"advanced\"\n" \
                                 "   id => \"$group\"\n" \
                                 "   add_field => { \"[@metadata][dataSource]\" => \"$group\" }\n" \
                                 "   type => \"azure\"\n" \
                                 "   decorate_events => true\n" \
                                 "   event_hubs => [\n" \
                                 "      $event_hubs\n" \
                                 "   ]\n" \
                                 " }\n "
                        return Template(azure).substitute(event_hubs=config, group=group)
                    else:
                        pass
                else:
                    return ""