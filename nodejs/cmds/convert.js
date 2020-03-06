const fs = require("fs")
const Minio = require("minio")
const { downloadJSON } = require("../src/download")
const { verifySlateData } = require("../src/verify")
const { convertToSlate } = require("../src/convert")
const { uploadFiles } = require("../src/upload")

exports.command = "convert [options]"
exports.describe = "## convert draft.js content to slate.js format ##"
exports.builder = yargs => {
  yargs
    .positional("minioHost", {
      type: "string",
      describe: "minio service host",
    })
    .positional("minioPort", {
      type: "number",
      describe: "minio service port",
    })
    .positional("accessKey", {
      type: "string",
      describe: "minio access key",
    })
    .positional("secretKey", {
      type: "string",
      describe: "minio secret key",
    })
    .positional("bucket", {
      type: "string",
      describe: "minio bucket to store content",
      default: "content",
    })
    .positional("userId", {
      type: "number",
      describe: "user ID to use for updating content",
    })
    .demandOption([
      "minioHost",
      "minioPort",
      "accessKey",
      "secretKey",
      "bucket",
      "userId",
    ])
    .help("h")
    .example(
      "convert --minioHost localhost --minioPort 9000 --accessKey foo --secretKey bar --userId 9999",
    )
}

exports.handler = async argv => {
  try {
    fs.mkdirSync("draftjs", { recursive: true })
    fs.mkdirSync("slate", { recursive: true })

    const minioClient = new Minio.Client({
      endPoint: argv.minioHost,
      port: argv.minioPort,
      useSSL: false,
      accessKey: argv.accessKey,
      secretKey: argv.secretKey,
    })

    const bucket = argv.bucket

    await downloadJSON(bucket, "draftjs", minioClient)
    await convertToSlate("draftjs", "slate", argv.userId)
    await verifySlateData("slate")
    await uploadFiles(bucket, "slate", minioClient)
  } catch (error) {
    console.log(error)
    process.exit(1)
  }
}
