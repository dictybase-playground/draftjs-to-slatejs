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

const downloadJSON = async (bucket, folder, minioClient) => {
  for (slug of slugs) {
    try {
      const filePath = `${folder}/${slug}.json`
      await minioClient.fGetObject(bucket, filePath, filePath)
      console.log(`downloaded ${slug}.json`)
    } catch (error) {
      console.log(error)
    }
  }
}

module.exports = {
  downloadJSON: downloadJSON,
}
