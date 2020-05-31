const crypto = require('crypto')


const PIC_B64_CHAR_EQUAL_REPLACE = '_';
const PIC_B64_CHAR_SLASH_REPLACE = '$';
const PIC_B64_CHAR_PLUS_REPLACE = '-';

const PIC_HASH_TYPE_LEN = 5;

const PIC_TYPE_BASE = 'b';
const PIC_TYPE_CROP = 'c';
const PIC_TYPE_HASH_BUILD = 'c';
const PIC_TYPE_SLIDE = 's';
const PIC_TYPE_MODEL = 'm';
const PIC_TYPE_BLUR = 'r';

const PIC_SIZE_BIG = 0;
const PIC_SIZE_NORMAL = 1;
const PIC_SIZE_SMALL = 2;
const PIC_SIZE_HASH_BUILD = 2;
const PIC_SIZE_VSMALL = 3;
const PIC_SIZE_VBIG_V3 = 4;
const PIC_SIZE_BIG_V3 = 5;

const FILE_EXT = '.jpg';

const PIC_ALL_TYPES = {
  base: PIC_TYPE_BASE,
  crop: PIC_TYPE_CROP,
  hash: PIC_TYPE_HASH_BUILD,
  slide: PIC_TYPE_SLIDE,
  model: PIC_TYPE_MODEL,
  blur: PIC_TYPE_BLUR,
};

const PIC_ALL_SIZE = {
  big: PIC_SIZE_BIG,
  normal: PIC_SIZE_NORMAL,
  small: PIC_SIZE_SMALL,
  hash: PIC_SIZE_HASH_BUILD,
  vsmall: PIC_SIZE_VSMALL,
  vbigv3: PIC_SIZE_VBIG_V3,
  bigv3: PIC_SIZE_BIG_V3,
};


function completeUrl(uuid, idGender, picId, dbHash) {
  return `${userPicturePath(uuid, idGender)}/${picturePath(picId, dbHash)}`;
}

function userPicturePath(uuid, idGenre) {
  const picDireactory = Math.floor(uuid / 1000);
  const b64 = Buffer.from(uuid)
                  .toString('base64')
                  .split('')
                  .reverse()
                  .join('')
                  .replace(/=/g, PIC_B64_CHAR_EQUAL_REPLACE)
                  .replace(/\+/g, PIC_B64_CHAR_PLUS_REPLACE)
                  .replace(/\//g, PIC_B64_CHAR_SLASH_REPLACE);

  return `${idGenre}/${picDireactory}/${b64}`;
}

function picturePath(uuid, dbHash) {
  const b64 = Buffer.from(uuid)
                  .toString('base64')
                  .replace(/=/g, PIC_B64_CHAR_EQUAL_REPLACE)
                  .replace(/\+/g, PIC_B64_CHAR_PLUS_REPLACE)
                  .replace(/\//g, PIC_B64_CHAR_SLASH_REPLACE)
                  .split('')
                  .reverse()
                  .join('');

  return `${b64}${dbHash}`;
}

function formatPath(uuid, type, size) {
  const data = `${uuid}${type}${size}`;
  const ddata = crypto.createHash('md5').update(data).digest('hex');
  return ddata.substring(0, PIC_HASH_TYPE_LEN);
}

const msg = {
  picture_id: '224747',
  picture_url:
      'https://apiv2.ltservices.ovh/filer/2/329/1gzM5IzM/3QzN0IjM9c5fc1383f94aa105a12bffdaf2a14701f19a05b91d2edc1c50802bc67de16010.jpg',
  picture_hash:
      '9c5fc1383f94aa105a12bffdaf2a14701f19a05b91d2edc1c50802bc67de1601',
  user_email: 'shainez29@mekmek.fr',
  user_id: '329385',
  user_gender: 2
}

const userPicPath =
    completeUrl(msg.user_id, msg.user_gender, msg.picture_id, msg.picture_hash);
console.log()
console.log(`${userPicPath}${formatPath(msg.user_id, 'b', 0)}${FILE_EXT}`)

    // https://apiv2.ltservices.ovh/filer/2/329/1gzM5IzM/2QzN0IjM0d6af850d0e20f9be6d4ecabc80e8cd3c553c6ba9159846918015b1364dafc0e81c04.jpg
    // https://apiv2.ltservices.ovh/filer/2/329/1gzM5IzM/2QzN0IjM0d6af850d0e20f9be6d4ecabc80e8cd3c553c6ba9159846918015b1364dafc0ee841e.jpg
    // https://apiv2.ltservices.ovh/filer/2/329/1gzM5IzM/2QzN0IjM0d6af850d0e20f9be6d4ecabc80e8cd3c553c6ba9159846918015b1364dafc0e0.jpg