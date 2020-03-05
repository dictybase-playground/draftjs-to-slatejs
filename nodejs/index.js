const fs = require("fs")
const { downloadJSON } = require("./src/download")
const { verifySlateData } = require("./src/verify")
const { convertToSlate } = require("./src/convert")
const { uploadFiles } = require("./src/upload")

const fullConversion = async () => {
  try {
    fs.mkdirSync("draftjs", { recursive: true })
    fs.mkdirSync("slate", { recursive: true })

    await downloadJSON("draftjs")
    await convertToSlate("draftjs", "slate")
    await verifySlateData("slate")
    await uploadFiles("slate")
  } catch (error) {
    console.log(error)
    process.exit(1)
  }
}

fullConversion()
