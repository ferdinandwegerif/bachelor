from json import decoder, load, dump


def main():
    sent = "Request sent(microseconds)"
    waiting = "Waiting(TTFB)(milliseconds)"
    download = "Content Download (milliseconds)"
    vugu_mean = {
        sent: 0,
        waiting: 0,
        download: 0
    }
    vue_mean = {
        sent: 0,
        waiting: 0,
        download: 0
    }
    with open("ajax_data.json", "r") as json_data:
        data = load(json_data)
        data = data["GET_EXAMPLE"]

        for framework, l in data.items():
            if framework == "Vugu":
                for data in l:
                    vugu_mean[sent] += float(data[sent])
                    vugu_mean[waiting] += float(data[waiting])
                    vugu_mean[download] += float(data[download])
                vugu_mean = divide_fields(vugu_mean, len(l))
            else:
                for data in l:
                    vue_mean[sent] += float(data[sent])
                    vue_mean[waiting] += float(data[waiting])
                    vue_mean[download] += float(data[download])
                vue_mean = divide_fields(vue_mean, len(l))
        json_data.close()
    print("VUGU MEAN VALUES: ", vugu_mean)
    print("Vue mean values: ", vue_mean)
    with open('results_get.json', 'w') as results:
        dump({'Vugu': vugu_mean, 'Vue': vue_mean}, results)
        results.close()


def divide_fields(values, divisor):
    new_val = values
    print(divisor)
    for k, v in new_val.items():
        new_val[k] = str(float(v/divisor))
    return new_val


if __name__ == '__main__':
    main()
