import * as ReactDOM from "react-dom";

export function render(component, containerID) {
    ReactDOM.render(component, document.getElementById(containerID))
}
