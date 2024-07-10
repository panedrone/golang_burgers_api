import React from "react";

import * as api from "./api";

import {StringField} from "./form_controls";
import fire from "./event_bus";
import {TableBurgers} from "./table_burgers";

let _burger_name = ""
let _ingredient_name = ""

let _burger_set = (_) => {
}
let _ingredient_set = (_) => {
}

fire.searchByIngredient = function (ingredient_name) {
    _burger_set("")
    _ingredient_set(ingredient_name)

    fetchSearchResults()
}

let _updateSearchResult = (_) => {
}

function fetchSearchResults() {
    let b = _burger_name !== null && _burger_name.length > 0;
    let i = _ingredient_name !== null && _ingredient_name.length > 0;
    if (!b && !i) {
        return
    }
    api.getJsonArray(`api/burgers/search?b=${_burger_name}&i=${_ingredient_name}`, (arr) => {
        _updateSearchResult(arr)
    })
}

export const TabSearch = () => {

    return <div>

        <table className="w100">
            <tbody>
            <tr>
                <td className="v-middle">
                    <StringField onChange={(v) => _burger_name = v} saveUpdater={(u) => _burger_set = u}
                                 placeholder={"Type a Search Key of Burger Name"}/>
                </td>
                <td className="v-middle w1">
                    <input type={"button"} value={"x"} onClick={() => _burger_set("")}/>
                </td>
                <td className="v-middle">
                    <StringField onChange={(v) => _ingredient_name = v} saveUpdater={(u) => _ingredient_set = u}
                                 placeholder={"Type an Search Key of Ingredient"}/>
                </td>
                <td className="v-middle w1">
                    <input type={"button"} value={"x"} onClick={() => _ingredient_set("")}/>
                </td>
            </tr>
            </tbody>
        </table>

        <button value={"Search"} onClick={() => fetchSearchResults()} className={""}>Search</button>

        <TableBurgers saveStateSetter={(s) => _updateSearchResult = s}/>

    </div>
}

