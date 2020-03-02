const fs = require("fs")
const draftToHtml = require("draftjs-to-html")
const convertToRaw = require("draft-js").convertToRaw
const convertFromRaw = require("draft-js").convertFromRaw

// Need to convert the raw JSON state to ContentState first
// then convert ContentState to a raw JS structure
// then finally convert that to HTML
// https://draftjs.org/docs/api-reference-data-conversion

const draftjsToHTML = (slug, content, folder) => {
  const contentState = convertFromRaw(JSON.parse(content))
  const raw = convertToRaw(contentState)
  const html = draftToHtml(raw)

  fs.writeFile(`${folder}/${slug}.html`, html, err => {
    if (err) {
      console.error(err)
    }
  })
}

module.exports = {
  draftjsToHTML: draftjsToHTML,
}
