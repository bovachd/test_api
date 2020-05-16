input = document.getElementById("input").innerHTML

async function callapi(event) {
    event.preventDefault()
    // console.log(input)
    
    fetch('http://localhost:8000/a', {
        method:'POST',
        headers:{
            'content-type' : 'application/json'
        },
        body: JSON.stringify({
            value: input.value,
        }),
    })
    .then(response => response.json())
    .then(data => console.log(data))
    // .then‌​(data => { console.log(data); })

    input.value = ""
}