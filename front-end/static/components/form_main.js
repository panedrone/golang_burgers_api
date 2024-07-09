import * as React from "react";

const ViewMode = Object.freeze({
    ABC: 1,
    SEARCH_BY_NAME: 2,
    SEARCH_BY_INGREDIENT: 3,
    MISSIONS_COMPLETED: 4,
});

export const Main = () => {

    const [viewMode, setViewMode] = React.useState(ViewMode.ABC);

    return <div>
        <table className="bg">
            <tbody>
            <tr>
                <td className="w1 nowrap v-middle">
                    <div className="card">
                        <a onClick={() => setViewMode(ViewMode.ABC)}>ABC</a>
                    </div>
                </td>
                <td className="w1 nowrap v-middle">
                    <div className="card">
                        <a onClick={() => setViewMode(ViewMode.SEARCH_BY_NAME)}>SEARCH BY NAME</a>
                    </div>
                </td>
                <td className="w1 nowrap v-middle">
                    <div className="card">
                        <a onClick={() => setViewMode(ViewMode.SEARCH_BY_INGREDIENT)}>SEARCH BY INGREDIENT</a>
                    </div>
                </td>
                <td className="w1 nowrap v-middle">
                    <div className="card">
                        <a onClick={() => setViewMode(ViewMode.MISSIONS_COMPLETED)}>Completed Missions</a>
                    </div>
                </td>
                <td>
                </td>
            </tr>
            </tbody>
        </table>
        <p>
        </p>
        <div>
            {
                viewMode === ViewMode.ABC
                    ?
                    <span>===ABC===</span>
                    :
                    viewMode === ViewMode.SEARCH_BY_NAME
                        ?
                        <span>===SEARCH_BY_NAME===</span>
                        :
                        viewMode === ViewMode.SEARCH_BY_INGREDIENT
                            ?
                            <span>===SEARCH_BY_INGREDIENT===</span>
                            :
                            <span>===???===</span>
            }
        </div>
    </div>
}