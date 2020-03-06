const fs = require("fs")
const path = require("path")
const { promisify } = require("util")
const draftToHtml = require("draftjs-to-html")
const { convertToRaw } = require("draft-js")
const { convertFromRaw } = require("draft-js")
const { html } = require("./deserialize")

const readdir = promisify(fs.readdir)
const readFile = promisify(fs.readFile)
const writeFile = promisify(fs.writeFile)

require("jsdom-global")()
global.DOMParser = window.DOMParser

const convertToSlate = async (inputFolder, outputFolder, userId) => {
  try {
    const files = await readdir(inputFolder)
    for (const file of files) {
      const fileContent = await readFile(`${inputFolder}/${file}`)
      // convert buffer object to json
      const json = JSON.parse(fileContent)
      // Need to convert the raw JSON state to ContentState first
      // then convert ContentState to a raw JS structure
      // then finally convert that to HTML
      // https://draftjs.org/docs/api-reference-data-conversion
      const contentState = convertFromRaw(
        JSON.parse(json.data.attributes.content),
      )
      const raw = convertToRaw(contentState)
      const convertedHTML = draftToHtml(raw)

      // now convert the HTML to Slate format
      const convertedSlateContent = html.deserialize(convertedHTML)
      const htmlString = JSON.stringify(convertedSlateContent)

      // JSON structure necessary for updating through content API
      const newJSON = {
        data: {
          type: "contents",
          id: json.data.id,
          attributes: {
            updated_by: userId,
            content: htmlString,
          },
        },
      }

      // output the new JSON files
      const filenameWithoutExtension = path.basename(file, path.extname(file))
      const newFile = `${outputFolder}/${filenameWithoutExtension}.json`
      await writeFile(newFile, JSON.stringify(newJSON))
      console.log(`✅  Successfully converted to ${newFile}`)
    }
  } catch (error) {
    console.log(error)
    process.exit(1)
  }
}

module.exports = {
  convertToSlate: convertToSlate,
}
