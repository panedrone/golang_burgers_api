import React from "react";

import * as api from "./api";
import {TableBurgers} from "./table_burgers";

const FirstLetters = () => {

    const [state, setState] = React.useState(null)

    function load() {
        api.getJson(`api/burgers/search/first-letters`, (json) => {
            setState(json)
        })
    }

    React.useEffect(() => {
        load()
    }, [])

    return <div>
        {
            state === null
                ?
                ""
                :
                <span>
                    {
                        state.map((letter, index) => {
                                return (
                                    <a key={index} style={{margin: "10px"}} href="#"
                                       onClick={() => loadByLetter(letter)}>{letter}</a>
                                )
                            }
                        )
                    }
                </span>
        }
    </div>
}

let loadByLetter = (letter) => {
    api.getJsonArray(`api/burgers/search/first-letter?letter=${letter}`, (arr) => {
        _loadList(arr)
    })
}

let _loadList = (_) => {

}

export const TabAbc = () => {

    return <div>
        <FirstLetters/>
        <TableBurgers saveStateSetter={(setter) => _loadList = setter}/>
    </div>
}