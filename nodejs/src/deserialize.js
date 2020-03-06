const Html = require("slate-html-serializer").default

const BLOCK_TAGS = {
  p: "paragraph",
  li: "list-item",
  ul: "bulleted-list",
  ol: "numbered-list",
  blockquote: "quote",
  pre: "code",
  h1: "heading-one",
  h2: "heading-two",
  h3: "heading-three",
  h4: "heading-four",
  h5: "heading-five",
  h6: "heading-six",
  figure: "figure",
  figcaption: "figcaption",
  hr: "divider",
  table: "table",
  th: "table-head",
  tr: "table-row",
  td: "table-cell",
  center: "center",
}

const MARK_TAGS = {
  strong: "bold",
  b: "bold",
  em: "italic",
  i: "italic",
  u: "underline",
  s: "strikethrough",
  code: "code",
}

const rules = [
  {
    deserialize(el, next) {
      const tagName = el.tagName.toLowerCase()
      // special case to grab href from links
      if (tagName === "a") {
        return {
          object: "inline",
          type: "link",
          nodes: next(el.childNodes),
          data: {
            href: el.getAttribute("href"),
          },
        }
      } else if (BLOCK_TAGS[tagName]) {
        return {
          object: "block",
          type: BLOCK_TAGS[tagName],
          nodes: next(el.childNodes),
        }
      } else if (MARK_TAGS[tagName]) {
        return {
          object: "mark",
          type: MARK_TAGS[tagName],
          nodes: next(el.childNodes),
        }
      }
    },
  },
]

// create a new serializer instance with our `rules` from above
const html = new Html({ rules: rules })

module.exports = {
  html: html,
}
