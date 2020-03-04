require("dotenv").config()
const fs = require("fs")
const { downloadJSONs } = require("./src/download")
const { draftjsToHTML } = require("./src/draft-to-html")
const { htmlToSlate } = require("./src/html-to-slate")
const { verifySlateData } = require("./src/verify")

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
    process.exit(1)
  }
}

fullConversion()
