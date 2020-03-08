

def extract_values(obj, key):
    #Prende tutti i valori della chiave passata dal JSON
    arr = []

    def extract(obj, arr, key):
        #Cerca ricorsivamente i valori della key nel JSON
        if isinstance(obj, dict):
            for k, v in obj.items():
                if isinstance(v, (dict, list)):     #se ho un dict anziche una list devo procedere diversamente e viceversa  
                    extract(v, arr, key)       #a questo punto se v è un dict o list, ripete il procedimento
                elif k == key:
                    arr.append(v)       #abbiamo trovato il valore, lo metto nell'array
        elif isinstance(obj, list):
            for item in obj:
                extract(item, arr, key)
        return arr

    results = extract(obj, arr, key)
    return results