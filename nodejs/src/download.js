const Minio = require("minio")

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

const minioClient = new Minio.Client({
  endPoint: "",
  //   port: 9000,
  useSSL: true,
  accessKey: "",
  secretKey: "",
})

const downloadJSON = async folder => {
  for (slug of slugs) {
    try {
      await minioClient.fGetObject(
        folder,
        `${folder}/${slug}.json`,
        `${folder}/${slug}.json`,
      )
      console.log(`downloaded ${slug}.json`)
    } catch (error) {
      console.log(error)
    }
  }
}

module.exports = {
  downloadJSON: downloadJSON,
}
