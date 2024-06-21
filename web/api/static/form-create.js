const err = document.querySelector("#err")
const btnRow = document.querySelector("#btnRow")
const btnAddField = document.querySelector("#btnAddField")
const btnSubmit = document.querySelector("#btnSubmit")

function setInputConstraints(inp, str) {
    switch (str) {
        case 'VARCHAR(128)':
        case 'VARCHAR(2048)':
        case 'TEXT':
        case 'MEDIUMTEXT':
        case 'LONGTEXT':
            inp.setAttribute("type", "text")
            break;
        case 'BOOL':
            inp.setAttribute("type", "checkbox")
            break;
        case 'INT':
            inp.setAttribute("type", "number")
            break;
        case 'FLOAT':
        case 'DATE':
        case 'TIME':
        case 'YEAR':
        case 'DATETIME':
        case 'TIMESTAMP':
            break;
        default:
            console.log(`Not implemented`);
    }
}


function addOpt(sel, str) {
    let o = document.createElement("option")
    o.setAttribute("value", str)
    o.text = str
    sel.appendChild(o)
}

// sel.addEventListener("change", (e) => {
//     setInputConstraints(inp, e.target.value)
// })


function createDataType(e) {
    e.preventDefault()
    let sel = document.createElement("select")
    let row = document.createElement("tr")

    addOpt(sel, "VARCHAR(128)")
    addOpt(sel, "VARCHAR(2048)")
    addOpt(sel, "TEXT")
    addOpt(sel, "MEDIUMTEXT")
    addOpt(sel, "LONGTEXT")
    addOpt(sel, "BOOL")
    addOpt(sel, "INT")
    addOpt(sel, "FLOAT")
    addOpt(sel, "DATE")
    addOpt(sel, "TIME")
    addOpt(sel, "YEAR")
    addOpt(sel, "DATETIME")
    addOpt(sel, "TIMESTAMP")

    let inp = document.createElement("input")
    inp.setAttribute("type", "text")
    inp.setAttribute("minlength", "1")

    inp.addEventListener("keyup", (e) => {
        sel.setAttribute("name", e.target.value)
    })

    let td1 = document.createElement("td")
    td1.appendChild(sel)

    let td2 = document.createElement("td")
    td2.appendChild(inp)

    row.append(td1, td2)


    btnRow.insertAdjacentElement("beforebegin", row)
}

btnAddField.addEventListener("click", createDataType)