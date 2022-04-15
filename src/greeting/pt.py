import sys

def ingredients(dish, serve):
    return {
        "dish": dish,
        "pork_belly": str(0.25*float(serve))+"kg",
        "dry_pea_pod": str(0.05*float(serve))+"kg",
    }
