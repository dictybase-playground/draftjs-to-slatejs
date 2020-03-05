const Minio = require("minio")

const minioClient = new Minio.Client({
  endPoint: "",
  //   port: 9000,
  accessKey: "",
  secretKey: "",
})

const uploadFiles = async () => {
  const exists = await minioClient.bucketExists("slate")
  if (!exists) {
    minioClient.makeBucket("slate", "us-east-1", err => {
      // if (err) return console.log(err)
      console.log('Bucket created successfully in "us-east-1".')
    })
  }
  //   minioClient.bucketExists("slate", (err, exists) => {
  //     if (err) {
  //       console.log(error)
  //     }
  //     if (!exists) {
  //       minioClient.makeBucket("slate", "us-east-1", err => {
  //         // if (err) return console.log(err)

  //         console.log('Bucket created successfully in "us-east-1".')
  //       })
  //     }
  //   })

  //   var metaData = {
  //     "Content-Type": "application/octet-stream",
  //     "X-Amz-Meta-Testing": 1234,
  //     example: 5678,
  //   }
  // Using fPutObject API upload your file to the bucket europetrip.
  //   minioClient.fPutObject(
  //     "slate",
  //     "about.json",
  //     "slate/dsc-deposit.json",
  //     function(err, etag) {
  //       if (err) return console.log(err)
  //       console.log("File uploaded successfully.")
  //     },
  //   )
}

uploadFiles()
