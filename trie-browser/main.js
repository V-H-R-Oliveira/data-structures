const trie = new tsTrie.Trie()

const loadData = async () => {
    const url = "https://jsonplaceholder.typicode.com/users"
    const request = await fetch(url)
    const data = await request.json()
    return data
}

const addSample = (parent, text) => {
    const li = document.createElement("li")
    const cleanText = text.trim()

    li.innerText = cleanText
    li.classList.add("sample")
    li.id = cleanText
    parent.appendChild(li)
}

const addSamplesToUI = async (data) => {
    const names = data.map(({ name }) => name)
    const sampleList = document.getElementById("sample-list")

    for (const name of names) {
        addSample(sampleList, name)
        trie.insert(name)
    }
}

const toggleSamplesVisibility = (matchText) => {
    const sampleList = document.getElementById("sample-list")

    for (const sampleElement of sampleList.children) {
        if (sampleElement.textContent.trim() != matchText) {
            console.log("Toggle", sampleElement.textContent)
            sampleElement.classList.toggle("hidden")
            continue
        }

        sampleElement.classList.toggle("top-order")
    }
}

window.onload = async () => {
    const data = await loadData()
    await addSamplesToUI(data)
}

// TODO: add interactive trie search method
(() => {
    const textInput = document.getElementById("text")
    const searchForm = document.getElementById("search-form")

    searchForm.addEventListener("submit", (event) => {
        event.preventDefault()

        const searchText = textInput.value.trim()
        const found = trie.search(searchText)

        if (found) {
            toggleSamplesVisibility(searchText)
            setTimeout(() => toggleSamplesVisibility(searchText), 1500)
        }

        textInput.value = ""
    })
})()