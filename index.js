require("dotenv").config()
const fs = require("fs")
const fetch = require("node-fetch")
const Value = require("slate").Value
const draftjsToHTML = require("./draft-to-html").draftjsToHTML
const htmlToSlate = require("./html-to-slate").htmlToSlate

const url = process.env.API_SERVER
const namespace = "dsc"
const slugs = [
  "intro",
  "about",
  "other-materials",
  "order",
  "payment",
  "deposit",
  "faq",
  "nomenclature-guidelines",
  "other-stock-centers",
]

const downloadJSONs = folder => {
  slugs.forEach(async item => {
    try {
      const slug = `${namespace}-${item}`
      const res = await fetch(`${url}/contents/slug/${slug}`)
      const json = await res.json()
      const content = json.data.attributes.content

      draftjsToHTML(slug, content, "html")
      fs.writeFile(`${folder}/${slug}.json`, JSON.stringify(json), err => {
        if (err) {
          console.error(err)
        }
      })
    } catch (error) {
      console.error(error)
    }
  })
}

const verifySlateJSONs = folder => {
  fs.readdir(folder, (err, files) => {
    if (err) {
      console.error(err)
      process.exit(1) // stop the script
    }
    files.forEach(file => {
      fs.readFile(file, "UTF-8", (err, content) => {
        const fileContent = fs.readFileSync(`${folder}/${file}`)
        try {
          const value = Value.fromJSON(JSON.parse(fileContent))
        } catch (error) {
          console.error(error)
        }
      })
    })
  })
}

const uploadJSONs = () => {}

fs.mkdirSync("json", { recursive: true })
fs.mkdirSync("html", { recursive: true })
fs.mkdirSync("slate", { recursive: true })
downloadJSONs("json")
htmlToSlate("html", "slate")
verifySlateJSONs("slate")
// for content PATCH requests, need:
// id, updated_by, content
// uploadJSONs()
