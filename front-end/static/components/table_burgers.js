import React, {useState} from "react";

import {IngredientList} from "./ingredient_list";

const BurgerRow = ({burger}) => {

    return <tr>
        <td className="w1">
            <img src={burger.image_url} style={{width: "320px", height: "auto", padding: "10px"}} alt={burger.name}/>
        </td>
        <td>
            <h3>{burger.name}</h3>
            <div>{burger.description}</div>
            <IngredientList initial={burger.ingredients}/>
        </td>
    </tr>
}

export const TableBurgers = ({saveStateSetter}) => {

    const [state, setState] = useState(null)

    if (saveStateSetter) {
        saveStateSetter(setState)
    }

    return <table className={"w100"}>
        {
            state === null
                ?
                <tbody></tbody>
                :
                <tbody>
                {
                    state.map((burger) => {
                            return (
                                <BurgerRow burger={burger}/>
                            )
                        }
                    )
                }
                </tbody>
        }
    </table>
}
