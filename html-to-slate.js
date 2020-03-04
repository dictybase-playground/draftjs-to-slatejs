const fs = require("fs")
const path = require("path")
const { promisify } = require("util")
const { html } = require("./deserialize")

require("jsdom-global")()
global.DOMParser = window.DOMParser

const readdir = promisify(fs.readdir)
const readFile = promisify(fs.readFile)
const writeFile = promisify(fs.writeFile)

const htmlToSlate = async (inputFolder, outputFolder) => {
  try {
    const files = await readdir(inputFolder)
    for (const file of files) {
      const fileContent = await readFile(`${inputFolder}/${file}`)
      const convertedHtml = html.deserialize(fileContent)
      const htmlString = JSON.stringify(convertedHtml)
      const filenameWithoutExtension = path.basename(file, path.extname(file))
      const newFile = `${outputFolder}/${filenameWithoutExtension}.json`

      await writeFile(newFile, htmlString)
      console.log(`✅  Successfully converted to ${newFile}`)
    }
  } catch (error) {
    console.log(error)
    process.exit(1)
  }
}

module.exports = {
  htmlToSlate: htmlToSlate,
}
