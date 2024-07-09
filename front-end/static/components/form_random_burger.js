import React, {useState} from "react";

import * as api from "./api"
import fire from "./event_bus";

let _burger = {
    "id": 19,
    "name": "Mojitto Burger",
    "description": "Burger a day keeps the fat away",
    "image_url": null,
    "ingredients": [
        {
            "id": 12,
            "name": "patty"
        },
        {
            "id": 82,
            "name": "bread"
        },
        {
            "id": 83,
            "name": "maida"
        },
        {
            "id": 84,
            "name": "chick"
        }
    ]
}

_burger = null

const IngredientList = ({initial}) => {
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
                                        <a href="#" onClick={() => fire.activateSearchByIngredient(ingredient.name)}>{ingredient.name}</a>
                                    </li>
                                )
                            }
                        )
                    }
                </ul>
        }
    </>
}

export const RandomBurger = () => {

    const [burgerData, setBurgerData] = useState(null)

    function randomBurger() {
        api.getJson(`api/burgers/random`, (json) => {
            setBurgerData(json)
        })
    }

    // https://stackoverflow.com/questions/53715465/can-i-set-state-inside-a-useeffect-hook
    // using setState inside useEffect will create an infinite loop
    // One of the popular cases that using useState inside of useEffect will not cause
    // an infinite loop is when you pass an empty array as a second argument
    // to useEffect like useEffect(() => {....}, []) which means
    // that the effect function should be called once: after the first mount/render only

    React.useEffect(() => {
        randomBurger()
    }, [])

    return <div>
        {
            burgerData === null ? "" :
                <div>
                    <div>
                        <table className={"bg"}>
                            <tbody>
                            <tr>
                                <td className="v-middle">
                                    <h2>{burgerData.name}</h2>
                                </td>
                                <td className="w1 v-middle">
                                    <a href="#">
                                        <input type="button" value="⟳" onClick={() => randomBurger()}/>
                                    </a>
                                </td>
                            </tr>
                            <tr>
                                <td colSpan={2}>
                                    <p>{burgerData.description}</p>
                                </td>
                            </tr>
                            </tbody>
                        </table>
                    </div>
                    <div>
                        <table>
                            <tbody>
                            <tr>
                                <td>
                                    {
                                        burgerData.image_url === null ? "" :
                                            <img src={burgerData.image_url} style={{width: "660px", height: "auto"}}/>
                                    }
                                </td>
                                <td>
                                    <IngredientList initial={burgerData.ingredients}/>
                                </td>
                            </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
        }
    </div>
}