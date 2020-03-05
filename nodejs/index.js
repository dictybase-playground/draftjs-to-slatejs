const fs = require("fs")
const { verifySlateData } = require("./src/verify")
const { convertToSlate } = require("./src/convert")

const fullConversion = async () => {
  try {
    fs.mkdirSync("draftjs", { recursive: true })
    fs.mkdirSync("slate", { recursive: true })

    await convertToSlate("draftjs", "slate")
    await verifySlateData("slate")
  } catch (error) {
    console.log(error)
    process.exit(1)
  }
}

fullConversion()
