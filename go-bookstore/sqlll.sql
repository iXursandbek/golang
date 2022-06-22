db.users.insert({
    name: 'Xursand',
    email: 'xose@mail.ru',
    age: 26,
    relationship: {
        name: 'Anna',
        age: 20
    },
    photos: ['photo1.jpg','img.png'],
    birthday: Date()
})

db.users.insertMany([
    {
        name: 'Xursand',
        email: 'xose@mail.ru',
        age: 26,
        photos: ['photo1.jpg','img.png'],
        birthday: Date()
    },
    {
        name: 'jimi',
        email: 'xosep@mail.ru',
        age: 23,
        photos: ['photo1.jpg','img.png'],
        birthday: Date()
    },
    {
        name: 'Eshmat',
        email: 'eshmat@mail.ru',
        age: 28,
        photos: ['photo1.jpg','imgame.png'],
        birthday: Date()
    }
])