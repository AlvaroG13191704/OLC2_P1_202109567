
async function request(endpoint: string, method: string, body: string = "") {
  return await fetch(endpoint, {
    method,
    headers: {
      "Content-Type": "text/plain"
    },
    body: body
  })
}



export async function analyzeAPI(code: string= "") {
  const response = await request("http://127.0.0.1:3000/analyze", "POST", code);


  return await response.json();
}