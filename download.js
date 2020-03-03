const fs = require("fs")
const fetch = require("node-fetch")

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

const downloadJSONs = async folder => {
  slugs.forEach(async item => {
    try {
      const slug = `${namespace}-${item}`
      const res = await fetch(`${url}/contents/slug/${slug}`)
      if (res.ok) {
        const json = await res.json()
        const content = json.data.attributes.content
        const file = `${folder}/${slug}.json`
        fs.writeFile(file, JSON.stringify(json), err => {
          if (err) {
            console.error(err)
          }
        })
        console.log(`✅  Successfully created ${file}`)
      } else {
        console.log(res.statusText)
        process.exit(1)
      }
    } catch (error) {
      console.log(error)
      process.exit(1)
    }
  })
}

module.exports = {
  downloadJSONs: downloadJSONs,
}
