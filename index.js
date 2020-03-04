require("dotenv").config()
const fs = require("fs")
const { downloadJSONs } = require("./download")
const { draftjsToHTML } = require("./draft-to-html")
const { htmlToSlate } = require("./html-to-slate")
const { verifySlateData } = require("./verify")

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
