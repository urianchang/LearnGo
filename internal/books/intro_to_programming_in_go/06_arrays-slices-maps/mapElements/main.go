package main

import "fmt"

func main() {
  // elements := make(map[string]string)
  //
  // elements["H"] = "Hydrogen"
  // elements["He"] = "Helium"
  // elements["Li"] = "Lithium"
  // elements["Be"] = "Beryllium"
  // elements["B"] = "Boron"
  // elements["C"] = "Carbon"
  // elements["N"] = "Nitrogen"
  // elements["O"] = "Oxygen"
  // elements["F"] = "Fluorine"
  // elements["Ne"] = "Neon"
  //
  // fmt.Println(elements["Li"])
  //
  // //: Returns zero value for key that doesn't exist
  // fmt.Println(elements["Um"])
  //
  // //: Check if element exists -> returns value, Boolean
  // name, ok := elements["Un"]
  // fmt.Println(name, ok)
  //
  // //: First try to get value from map, then if successful, run code block
  // if name, ok := elements["Um"]; ok {
  //   fmt.Println(name, ok)
  // }

  //: Faster ways of creating maps
  elements := map[string]map[string]string{
    "H": map[string]string{
      "name": "Hydrogen",
      "state": "gas",
    },
    "He": map[string]string{
      "name": "Helium",
      "state": "gas",
    },
    "Li": map[string]string{
      "name": "Lithium",
      "state": "solid",
    },
    "Be": map[string]string{
      "name": "Beryllium",
      "state": "solid",
    },
    "B": map[string]string{
      "name": "Boron",
      "state": "solid",
    },
    "C": map[string]string{
      "name": "Carbon",
      "state": "solid",
    },
    "N": map[string]string{
      "name": "Nitrogen",
      "state": "gas",
    },
    "O": map[string]string{
      "name": "Oxygen",
      "state": "gas",
    },
    "F": map[string]string{
      "name": "Fluorine",
      "state": "gas",
    },
    "Ne": map[string]string{
      "name": "Neon",
      "state": "gas",
    },
  }

  /* ====
    Accessing an element of a map can return two values instead of just one.
    The first value is the result of the lookup, the second tells us whether
    or not the lookup was successful.
  ==== */
  
  if el, ok := elements["Li"]; ok {
    fmt.Println(el["name"], el["state"])
  }

}
