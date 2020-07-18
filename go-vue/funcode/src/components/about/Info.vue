<template>
    <v-flex xs12 sm6 md12>
        <div v-if="status">
            <v-progress-circular :size="50" color="primary" indeterminate>
            </v-progress-circular>
        </div>
        <span v-if="!status">Total Perusahan IT yang terdaftar {{counterLinks}} </span>
        <br>
        <span v-if="msg != '' && counterLinks == 0">{{msg}}</span>
        <v-spacer></v-spacer>
        <a @click="removeAllData" style="font-weight:bold;">Remove All Data</a>
    </v-flex>
</template>

<script>
import { mapGetters, mapActions, mapMutations } from "vuex";
export default {
    name: "Info",
    data (){
        return {
            status: false,
            msg: ''
        }
    },
    computed: {
        ...mapGetters([
            'counterLinks',
        ])
    },
    methods: {
        ...mapMutations([
           'REMOVE_All'
        ]),
        ...mapActions([
            'removeAll'
        ]),
        removeAllData: function() {
            this.status = true;
            this.removeAll().then(() => {
                this.msg = "Semua data telah dihapus";
                this.status = false;
            })
        }
    }
}
</script>
