import React from "react";

import fire from "./event_bus";

export const IngredientList = ({initial}) => {
    return <>
        {
            initial === null
                ?
                ""
                :
                <ul>
                    {
                        initial.map((ingredient, index) => {
                                return (
                                    <li key={index} className="nowrap">
                                        <a href="#" onClick={() => fire.activateSearchByIngredient(ingredient.name)}>
                                            {ingredient.name}
                                        </a>
                                    </li>
                                )
                            }
                        )
                    }
                </ul>
        }
    </>
}

