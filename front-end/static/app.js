import React from "react";

import * as shared from "./components/shared";
import fire from './components/event_bus'
import {ErrorArea} from "./components/error_area";

function _assignServerErrorUpdater(updater) {
    fire.showServerError = updater
}

const App = () => {

    React.useEffect(() => {
        windowOnLoad();
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
                                                   href="https://github.com/panedrone/golang_assessment_spy_cat_agency">GitHub
                                                    Repository</a>
                                            </div>
                                        </td>
                                        <td className="w1 nowrap v-middle">
                                            <div className="card">
                                                <a target="_blank" href="swagger/index.html">API Docs</a>
                                            </div>
                                        </td>
                                    </tr>
                                    </tbody>
                                </table>
                            </td>
                            <td className="w1 v-middle">
                                <div className="card">
                                    {
                                        viewMode === ViewMode.LOGIN ?
                                            <span>Welcome!</span>
                                            :
                                            <a onClick={() => setViewMode(ViewMode.LOGIN)} href="#">
                                                Logout
                                            </a>
                                    }
                                </div>
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
                    {
                        viewMode === ViewMode.LOGIN
                            ?
                            <LoginForm/>
                            :
                            viewMode === ViewMode.AGENCY
                                ?
                                <Agency/>
                                :
                                <SpyCat initialSpyCat={_currentSpyCat}/>
                    }
                </div>
            </td>
        </tr>
        </tbody>
    </table>;
}

shared.render(<App/>, "app")

function windowOnLoad() {
    whoiam.fetchWhoIAm()
}

// windowOnLoad().then(() => console.log('== windowOnLoad() completed =='))