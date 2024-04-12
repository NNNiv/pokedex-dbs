/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./internal/templates/**/*.templ"],
  theme: {
    colors: {
      pkmnbg: "#35374B",
      pkmnfg: "#7DD696",
      pkmnfg2: "#1B2729",
      search: "#888DC6",
      searchbtn: "#D9D9D9",
    },
    fontFamily: {
      pkmn: ["PKMN RBYGSC"],
    },
  },
};
