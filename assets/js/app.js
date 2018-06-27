

new Vue({
    el: 'body',

    data: {
        selected: "Pet store",
        owners: [],
        pets: [],
        newPet: {}
    },

    // This is run whenever the page is loaded to make sure we have a current pet list
    created: function() {
        this.$http.get('/pets').then(function(response) {
            this.pets = response.data.items? response.data.items: []
        })

        this.$http.get('/owners').then(function(response) {
            this.owners = response.data.owners? response.data.owners: []
        })
    },

    methods: {
        createPet: function() {
            if (!$.trim(this.newPet.name) && !$.trim(this.newPet.type)) {
                this.newPet = {}
                return
            }

             // Post the new pet to the /pets route using the $http client
            this.$http.put('/pets', this.newPet).success(function(response) {
                this.newPet.id = response.created
                this.pets.push(this.newPet)
                console.log("Pet created!")
                console.log(this.newPet)
                this.newPet = {}
            }).error(function(error) {
                console.log(error)
            });
        },

        // Use the $http client to delete a pet by its id
        deletePet: function(index) {
            this.$http.delete('/pets/' + this.pets[index].id).success(function(response) {
                this.pets.splice(index, 1)
                console.log("Pet deleted!")
            }).error(function(error) {
                console.log(error)
            })
        }
    }
})