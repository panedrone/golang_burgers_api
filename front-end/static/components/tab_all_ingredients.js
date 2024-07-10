import React from "react";

import * as api from "./api";
import fire from "./event_bus";

export const TabAllIngredients = () => {

    const [data, setData] = React.useState([])

    function _fetchAll() {
        api.getJsonArray(`api/ingredients`, (arr) => {
            setData(arr)
        })
    }

    React.useEffect(() => {
        _fetchAll()
    }, [])

    return <>
        <h3>Ingredients</h3>
        <table className="section">
            <thead>
            <tr>
                <th className="nowrap">ID</th>
                <th className="w1 nowrap center">Name</th>
            </tr>
            </thead>
            <tbody>
            {
                data.map((ingredient, index) => {
                        return (
                            <tr key={index}>
                                <td className="center w1">
                                    {ingredient.id}
                                </td>
                                <td>
                                    <a href="#" onClick={() => fire.activateSearchByIngredient(ingredient.name)}>
                                        {ingredient.name}
                                    </a>
                                </td>
                            </tr>
                        )
                    }
                )
            }
            </tbody>
        </table>
    </>
}
