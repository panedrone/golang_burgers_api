import * as React from "react";
import {RandomBurger} from "./form_random_burger";
import fire from "./event_bus";

const ViewMode = Object.freeze({
    RANDOM: 1,
    ABC: 2,
    SEARCH_BY_NAME: 3,
    SEARCH_BY_INGREDIENT: 4,
    MISSIONS_ALL_INGREDIENTS: 5,
    CREATE_BURGER: 6,
});

fire.activateSearchByIngredient = (ingredient_name) => {
    _setViewMode(ViewMode.SEARCH_BY_INGREDIENT)
    fire.searchByIngredient(ingredient_name)
}

let _setViewMode = (_) => {
}

export const Main = () => {

    const [viewMode, setViewMode] = React.useState(ViewMode.RANDOM);

    _setViewMode = setViewMode

    return <div>
        <table className="bg">
            <tbody>
            <tr>
                <td className="w1 nowrap v-middle">
                    <div className="card">
                        <a onClick={() => setViewMode(ViewMode.RANDOM)}>Random Burger</a>
                    </div>
                </td>
                <td className="w1 nowrap v-middle">
                    <div className="card">
                        <a onClick={() => setViewMode(ViewMode.ABC)}>ABC</a>
                    </div>
                </td>
                <td className="w1 nowrap v-middle">
                    <div className="card">
                        <a onClick={() => setViewMode(ViewMode.SEARCH_BY_NAME)}>Search by Burger Name</a>
                    </div>
                </td>
                <td className="w1 nowrap v-middle">
                    <div className="card">
                        <a onClick={() => setViewMode(ViewMode.SEARCH_BY_INGREDIENT)}>Search by Ingredient</a>
                    </div>
                </td>
                <td className="w1 nowrap v-middle">
                    <div className="card">
                        <a onClick={() => setViewMode(ViewMode.MISSIONS_ALL_INGREDIENTS)}>Ingredients</a>
                    </div>
                </td>
                <td className="w1 nowrap v-middle">
                    <div className="card">
                        <a onClick={() => setViewMode(ViewMode.CREATE_BURGER)}>New Burger</a>
                    </div>
                </td>
                <td>
                </td>
            </tr>
            </tbody>
        </table>
        <p>
        </p>
        <div className={"card"}>
            {
                viewMode === ViewMode.RANDOM
                    ?
                    <RandomBurger/>
                    :
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