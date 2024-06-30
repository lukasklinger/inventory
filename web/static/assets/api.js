// file captures common functions used to interact with the API

// fetch all entries from the server, store in global variable for future processing
async function fetchEntries() {
    try {
        let response = await fetch("./api/entry")
        if(!response.ok) {
            throw new Error (`Response status when fetching entries: "${response.status}`)
        }

        entries = await response.json()
    } catch (error) {
        alert(error)
    }
}