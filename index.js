require("dotenv").config()
const fs = require("fs")
const downloadJSONs = require("./download").downloadJSONs
const draftjsToHTML = require("./draft-to-html").draftjsToHTML
const htmlToSlate = require("./html-to-slate").htmlToSlate
const verifySlateData = require("./verify").verifySlateData

// still needs to be fixed to run synchronously
const fullConversion = async () => {
  try {
    fs.mkdirSync("draftjs", { recursive: true })
    fs.mkdirSync("html", { recursive: true })
    fs.mkdirSync("slate", { recursive: true })

    await downloadJSONs("draftjs")
    await draftjsToHTML("draftjs", "html")
    await htmlToSlate("html", "slate")
    await verifySlateData("slate")
  } catch (error) {
    console.log(error)
  }
}

fullConversion()
