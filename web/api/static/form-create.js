const err = document.querySelector("#err")
const actionRow = document.querySelector("#actionRow")
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


function opt(sel, str) {
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
    opt(sel, "VARCHAR(128)")
    opt(sel, "VARCHAR(2048)")
    opt(sel, "TEXT")
    opt(sel, "MEDIUMTEXT")
    opt(sel, "LONGTEXT")
    opt(sel, "BOOL")
    opt(sel, "INT")
    opt(sel, "FLOAT")
    opt(sel, "DATE")
    opt(sel, "TIME")
    opt(sel, "YEAR")
    opt(sel, "DATETIME")
    opt(sel, "TIMESTAMP")

    let inp = document.createElement("input")
    inp.setAttribute("type", "text")
    inp.setAttribute("name", "VARCHAR(128)")
    inp.setAttribute("minlength", "1")

    sel.addEventListener("change", (e) => {
        inp.setAttribute("name", e.target.value)
    })

    let td1 = document.createElement("td")
    td1.appendChild(sel)

    let td2 = document.createElement("td")
    td2.appendChild(inp)

    let row = document.createElement("tr")
    row.append(td1, td2)


    actionRow.insertAdjacentElement("beforebegin", row)
}

btnAddField.addEventListener("click", createDataType)
