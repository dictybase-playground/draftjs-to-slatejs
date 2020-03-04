const fs = require("fs")
const path = require("path")
const { promisify } = require("util")
const draftToHtml = require("draftjs-to-html")
const { convertToRaw } = require("draft-js")
const { convertFromRaw } = require("draft-js")

const readdir = promisify(fs.readdir)
const readFile = promisify(fs.readFile)
const writeFile = promisify(fs.writeFile)

// Need to convert the raw JSON state to ContentState first
// then convert ContentState to a raw JS structure
// then finally convert that to HTML
// https://draftjs.org/docs/api-reference-data-conversion

const draftjsToHTML = async (inputFolder, outputFolder) => {
  try {
    const files = await readdir(inputFolder)
    for (const file of files) {
      const fileContent = await readFile(`${inputFolder}/${file}`)
      // convert buffer object to json
      const json = JSON.parse(fileContent)
      const contentState = convertFromRaw(
        JSON.parse(json.data.attributes.content),
      )
      const raw = convertToRaw(contentState)
      const html = draftToHtml(raw)
      const filenameWithoutExtension = path.basename(file, path.extname(file))

      const newFile = `${outputFolder}/${filenameWithoutExtension}.html`
      await writeFile(newFile, html)
      console.log(`✅  Successfully converted to ${newFile}`)
    }
  } catch (error) {
    console.log(error)
    process.exit(1)
  }
}

module.exports = {
  draftjsToHTML: draftjsToHTML,
}
