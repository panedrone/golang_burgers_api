import React from "react";

import * as shared from "./components/shared";
import fire from './components/event_bus'
import {ErrorArea} from "./components/error_area";
import {Main} from "./components/form_main";

function _assignServerErrorUpdater(updater) {
    fire.showServerError = updater
}

const App = () => {

    React.useEffect(() => {
        //
    });

    return <table className="bg">
        <tbody>
        <tr>
            <td>
                <div className="card">
                    <table className="bg">
                        <tbody>
                        <tr>
                            <td className="v-middle">
                                <h2 id="whoiam">Burgers API Demo. React.js {React.version}</h2>
                            </td>
                            <td className="w1 v-middle">

                                <table>
                                    <tbody>
                                    <tr>
                                        <td className="w1 nowrap v-middle">
                                            <div className="card">
                                                <a target="_blank"
                                                   href="https://github.com/panedrone/golang_burgers_api">GitHub Repo</a>
                                            </div>
                                        </td>
                                        <td className="w1 nowrap v-middle">
                                            <div className="card">
                                                <a target="_blank" href="swagger/index.html">Open API Docs</a>
                                            </div>
                                        </td>
                                    </tr>
                                    </tbody>
                                </table>
                            </td>
                        </tr>
                        </tbody>
                    </table>
                </div>
            </td>
        </tr>
        <tr>
            <td>
                <div className="red-banner">
                    <ErrorArea saveUpdater={_assignServerErrorUpdater}/>
                </div>
            </td>
        </tr>
        <tr>
            <td>
                <div>
                    <Main/>
                </div>
            </td>
        </tr>
        </tbody>
    </table>;
}

shared.render(<App/>, "app")

// function windowOnLoad() {
//     whoiam.fetchWhoIAm()
// }

// windowOnLoad().then(() => console.log('== windowOnLoad() completed =='))