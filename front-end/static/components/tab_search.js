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
        _hideLogo()
        _updateSearchResult(arr)
    })
}

let _hideLogo = () => {}

export const TabSearch = () => {

    const ref_logo = React.useRef()

    _hideLogo = () => {
        ref_logo.current.style.display = "none"
    }

    return <div>

        <table className="w100 edit-form">
            <tbody>
            <tr>
                <td className="v-middle w1 nowrap">
                    Burger Name
                </td>
                <td className="v-middle">
                    <StringField onChange={(v) => _burger_name = v} saveUpdater={(u) => _burger_set = u}
                                 placeholder={"Type a Search Key"}/>
                </td>
                <td className="v-middle w1">
                    <input type={"button"} value={"x"} onClick={() => _burger_set("")}/>
                </td>
                <td className="v-middle w1 nowrap">
                    Ingredient
                </td>
                <td className="v-middle">
                    <StringField onChange={(v) => _ingredient_name = v} saveUpdater={(u) => _ingredient_set = u}
                                 placeholder={"Type an Search Key"}/>
                </td>
                <td className="v-middle w1">
                    <input type={"button"} value={"x"} onClick={() => _ingredient_set("")}/>
                </td>
                <td>
                </td>
                <td className="v-middle">
                    <input type={"button"} value={"Search"} onClick={() => fetchSearchResults()}/>
                </td>
            </tr>
            </tbody>
        </table>

        <div ref={ref_logo} style={{
            minHeight: "320px",
            backgroundRepeat: "no-repeat",
            backgroundPosition: "center center",
            backgroundImage: "url(static/search.png)"
        }}>
        </div>

        <TableBurgers saveStateSetter={(s) => _updateSearchResult = s}/>


    </div>
}

