import * as React from "react";
import fire from "./event_bus";
import {TabRandomBurger} from "./tab_random_burger";
import {TabAllIngredients} from "./tab_all_ingredients";
import {TabAbc} from "./tab_abc";
import {TabSearch} from "./tab_search";
import {TabNewBurger} from "./tab_new_burger";

fire.activateSearchByIngredient = (ingredient_name) => {
    _showTabSearch()
    fire.searchByIngredient(ingredient_name)
}

let _showTabSearch = () => {
}

export const Main = () => {

    const ref_random = React.useRef()
    const ref_abs = React.useRef()
    const ref_search = React.useRef()
    const ref_all_ingredients = React.useRef()
    const ref_new_burger = React.useRef()

    let ref_active = ref_random

    _showTabSearch = function () {
        show(ref_search)
    }

    function show(ref) {
        hide(ref_active)
        ref.current.style.display = "block"
        ref_active = ref
    }

    function hide(ref) {
        ref.current.style.display = "none";
    }

    const hidden = {display: "none"}

    return <div>
        <table className="bg">
            <tbody>
            <tr>
                <td className="w1 nowrap v-middle">
                    <div className="card">
                        <a onClick={() => show(ref_abs)}>ABC</a>
                    </div>
                </td>
                <td className="w1 nowrap v-middle">
                    <div className="card">
                        <a onClick={() => show(ref_search)}>Search</a>
                    </div>
                </td>
                <td className="w1 nowrap v-middle">
                    <div className="card">
                        <a onClick={() => show(ref_all_ingredients)}>Ingredients</a>
                    </div>
                </td>
                <td className="w1 nowrap v-middle">
                    <div className="card">
                        <a onClick={() => show(ref_random)}>Random Burger</a>
                    </div>
                </td>
                <td className="w1 nowrap v-middle">
                    <div className="card">
                        <a onClick={() => show(ref_new_burger)}>New Burger</a>
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
            <div ref={ref_abs} style={hidden}>
                <TabAbc/>
            </div>
            <div ref={ref_search} style={hidden}>
                <TabSearch/>
            </div>
            <div ref={ref_all_ingredients} style={hidden}>
                <TabAllIngredients/>
            </div>
            <div ref={ref_random}>
                <TabRandomBurger/>
            </div>
            <div ref={ref_new_burger} style={hidden}>
                <TabNewBurger/>
            </div>
        </div>
    </div>
}