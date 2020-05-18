const app = require("express")()
app.get("/", (req, res) => res.send("hi"))
app.listen(process.env.OPTIC_API_PORT || 3000)
