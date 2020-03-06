const slugs = [
  "dsc-intro",
  "dsc-about",
  "dsc-other-materials",
  "dsc-order",
  "dsc-payment",
  "dsc-deposit",
  "dsc-faq",
  "dsc-nomenclature-guidelines",
  "dsc-other-stock-centers",
]

const downloadJSON = async (folder, minioClient) => {
  for (slug of slugs) {
    try {
      const filePath = `${folder}/${slug}.json`
      await minioClient.fGetObject(folder, filePath, filePath)
      console.log(`downloaded ${slug}.json`)
    } catch (error) {
      console.log(error)
    }
  }
}

module.exports = {
  downloadJSON: downloadJSON,
}
