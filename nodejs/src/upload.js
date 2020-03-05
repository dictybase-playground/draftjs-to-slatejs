const fs = require("fs")
const path = require("path")
const { promisify } = require("util")
const Minio = require("minio")

const readdir = promisify(fs.readdir)

const minioClient = new Minio.Client({
  endPoint: "",
  //   port: 9000,
  useSSL: true,
  accessKey: "",
  secretKey: "",
})

const uploadFiles = async folder => {
  try {
    const files = await readdir(folder)
    for (const file of files) {
      const fileName = `${folder}/${file}`
      const metaData = {
        "Content-Type": "application/json",
      }
      await minioClient.fPutObject("draftjs", fileName, fileName, metaData)
      console.log(`Successfully uploaded ${file}`)
    }
  } catch (error) {
    console.error(error)
    process.exit(1)
  }
}

module.exports = {
  uploadFiles: uploadFiles,
}
