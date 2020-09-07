import "./scss/agent.scss"
import "./bootstrap";

import Agent from "./views/index";
import AgentForm from "./views/agentform";

const components = {
    'agent': Agent,
    'agent-form': AgentForm,
};

const install = function (Vue, options) {
    if (install.installed) return;
    Object.keys(components).forEach(key => Vue.component(key, components[key]));
};

if (typeof window !== 'undefined' && window.Vue) {
    install(window.Vue);
}

export default install
