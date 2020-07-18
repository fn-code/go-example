<template>

        <v-layout row wrap>
             <v-flex xs12 sm12>
         <v-card flat style="padding-top:20px;">
            <v-card-text>
              <v-avatar
                size="88"
              >
                <v-img
                  :src="`https://secure.gravatar.com/avatar/0bd6cc64aafc39299b5b46c31956f4af.jpg?s=80&r=g&d=mm`"
                  class="mb-4"
                ></v-img>
              </v-avatar>

              <h3 class="headline mb-2">
                {{ userDetail.name }}
              </h3>
              <div class="blue--text mb-2">{{ userDetail.email }}</div>
              <div class="blue--text subheading font-weight-bold">{{ userDetail.username }}</div>
            </v-card-text>
            <v-divider></v-divider>
            <v-flex sm6>
            <v-layout
              tag="v-card-text"
              text-xs-left
              wrap>

              <v-flex tag="strong" xs5 text-xs-right mr-3 mb-2>Company:</v-flex>
              <v-flex>{{ userDetail.company.name }}</v-flex>
              <v-flex tag="strong" xs5 text-xs-right mr-3 mb-2>Website:</v-flex>
              <v-flex>
                <a :href="`//${userDetail.website}`" target="_blank">{{ userDetail.website }}</a>
              </v-flex>
              <v-flex tag="strong" xs5 text-xs-right mr-3 mb-2>Phone:</v-flex>
              <v-flex>{{ userDetail.phone }}</v-flex>
            </v-layout>
            </v-flex>


          </v-card>
             </v-flex>
        </v-layout>
</template>

<script>
import { mapMutations, mapState } from 'vuex';
const avatars = [
    '?accessoriesType=Blank&avatarStyle=Circle&clotheColor=PastelGreen&clotheType=ShirtScoopNeck&eyeType=Wink&eyebrowType=UnibrowNatural&facialHairColor=Black&facialHairType=MoustacheMagnum&hairColor=Platinum&mouthType=Concerned&skinColor=Tanned&topType=Turban',
    '?accessoriesType=Sunglasses&avatarStyle=Circle&clotheColor=Gray02&clotheType=ShirtScoopNeck&eyeType=EyeRoll&eyebrowType=RaisedExcited&facialHairColor=Red&facialHairType=BeardMagestic&hairColor=Red&hatColor=White&mouthType=Twinkle&skinColor=DarkBrown&topType=LongHairBun',
    '?accessoriesType=Prescription02&avatarStyle=Circle&clotheColor=Black&clotheType=ShirtVNeck&eyeType=Surprised&eyebrowType=Angry&facialHairColor=Blonde&facialHairType=Blank&hairColor=Blonde&hatColor=PastelOrange&mouthType=Smile&skinColor=Black&topType=LongHairNotTooLong',
    '?accessoriesType=Round&avatarStyle=Circle&clotheColor=PastelOrange&clotheType=Overall&eyeType=Close&eyebrowType=AngryNatural&facialHairColor=Blonde&facialHairType=Blank&graphicType=Pizza&hairColor=Black&hatColor=PastelBlue&mouthType=Serious&skinColor=Light&topType=LongHairBigHair',
    '?accessoriesType=Kurt&avatarStyle=Circle&clotheColor=Gray01&clotheType=BlazerShirt&eyeType=Surprised&eyebrowType=Default&facialHairColor=Red&facialHairType=Blank&graphicType=Selena&hairColor=Red&hatColor=Blue02&mouthType=Twinkle&skinColor=Pale&topType=LongHairCurly'
  ]

export default {
    name: 'DetailPengguna',
    data() {
        return {
            idUsers: this.$route.params.id,
            avatar: null,
        };
    },
    created: function () {
    this.getUser(this.$route.params.id)
  },
  computed: {
      ...mapState(['userDetail'])
  },
   watch: {
      selected: 'randomAvatar'
    },

  methods: {
      ...mapMutations(['fetchUsersByID']),
      getUser:function(id) {
          this.fetchUsersByID(id)
      },
      randomAvatar () {
        this.avatar = avatars[Math.floor(Math.random() * avatars.length)]
      }
  }
}
</script>
