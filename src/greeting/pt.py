import sys

def ingredients(dish, serve):
    return {
        "dish": dish,
        "pork_belly": str(0.25*serve)+"kg",
        "dry_pea_pod": str(0.05*serve)+"kg",
    }
