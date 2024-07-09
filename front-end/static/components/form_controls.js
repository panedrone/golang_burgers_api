import * as React from "react";

export const PasswordField = ({onChange, saveUpdater}) => {

    const refDom = React.useRef()

    function _setValue(value) {
        if (onChange) {
            onChange(value)
        }
        refDom.current.value = value
    }

    if (saveUpdater) {
        saveUpdater((value) => _setValue(value))
    }

    function handleChange(event) {
        _setValue(event.target.value)
    }

    return (
        <label>
            <input type="password" ref={refDom} onChange={handleChange}/>
        </label>
    )
}

export const StringField = ({onChange, saveUpdater, initial=""}) => {

    const refDom = React.useRef()

    function _setValue(value) {
        if (onChange) {
            onChange(value)
        }
        refDom.current.value = value
    }

    if (saveUpdater) {
        saveUpdater((value) => _setValue(value))
    }

    React.useEffect(() => {
        refDom.current.value = initial
    })

    function handleChange(event) {
        _setValue(event.target.value)
    }

    return (
        <label>
            <input type="text" ref={refDom} onChange={handleChange}/>
        </label>
    )
}

export const IntegerField = ({onChange, saveUpdater, min = 1, max = 10}) => {

    const refDom = React.useRef()
    const refValue = React.useRef(min)

    function _setValue(value) {
        let parsed = parseInt(value.toString())
        if (parsed < min || parsed > max) {
            parsed = refValue.current // === panedrone: reset to the old one
        }
        refValue.current = parsed
        value = parsed.toString()
        if (onChange) {
            onChange(value)
        }
        refDom.current.value = value
    }

    if (saveUpdater) {
        saveUpdater(_setValue)
    }

    function handleChange(event) {
        let value = event.target.value
        if (!value) {
            return // === panedrone: allow typing from scratch
        }
        _setValue(value);
    }

    function handleUp() {
        let current = refValue.current
        if (current >= max) {
            return
        }
        current = current + 1
        _setValue(current);
    }

    function handleDown() {
        let current = refValue.current
        if (current <= min) {
            return
        }
        current = current - 1
        _setValue(current);
    }

    // === panedrone: <input type="number" is buggy:
    //      - impossible to disable typing of not-numbers
    //      - with invalid values, "onChange" is not fired

    // return (
    //     <label>
    //         <input type="number" min="1" max="10" pattern="[0-9\s]" value={value)} onChange={handleChange}/>
    //     </label>
    // )

    return (
        <div style={{display: "flex"}}>
            <input type="text" min={min} max={max} pattern="[0-9\s]"
                   ref={refDom} placeholder="Type a number..." onChange={handleChange} style={{flex: 1}}/>
            <div>
                <button style={{padding: 0, margin: 1, display: "block", lineHeight: "0.7em"}} onClick={handleUp}>
                    &#x25B4;
                </button>
                <button style={{padding: 0, margin: 1, display: "block", lineHeight: "0.7em"}} onClick={handleDown}>
                    &#x25BE;
                </button>
            </div>
        </div>
    )
}

export const TextAreaField = ({onChange, saveUpdater, cols=40, rows=5, initial=""}) => {

    const refDom1 = React.useRef()

    function _setValue(value) {
        if (onChange) {
            onChange(value)
        }
        refDom1.current.value = value
    }

    if (saveUpdater) {
        saveUpdater((value) => _setValue(value))
    }

    React.useEffect(() => {
        refDom1.current.value = initial
    })

    function handleChange(event) {
        _setValue(event.target.value)
    }

    return (
        <label>
            <textarea cols={cols} rows={rows} ref={refDom1} onChange={handleChange}></textarea>
        </label>
    )
}
