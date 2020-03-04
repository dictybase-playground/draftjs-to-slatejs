const fs = require("fs")
const { promisify } = require("util")
const { Value } = require("slate")

const readdir = promisify(fs.readdir)
const readFile = promisify(fs.readFile)

const verifySlateData = async folder => {
  try {
    const files = await readdir(folder)
    for (const file of files) {
      const fileName = `${folder}/${file}`
      const fileContent = await readFile(fileName)
      Value.fromJSON(JSON.parse(fileContent))
      console.log(`✅  ${fileName} is valid Slate data`)
    }
  } catch (error) {
    console.error(error)
    process.exit(1)
  }
}

module.exports = {
  verifySlateData: verifySlateData,
}
