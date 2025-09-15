import os
import sys
import random
import time


def check_args(file_args):

    try:
        if len(file_args) != 2 or int(file_args[1]) <= 0:
            raise Exception()
    except:
        print("Usage:  create_measurements.sh <positive integer number of records to create>")
        print("        You can use underscore notation for large number of records.")
        print("        For example:  1_000_000_000 for one billion")
        exit()


def build_weather_station_name_list():

    station_names = []
    with open('../../../data/weather_stations.csv', 'r') as file:
        file_contents = file.read()
    for station in file_contents.splitlines():
        if "#" in station:
            next
        else:
            station_names.append(station.split(';')[0])
    return list(set(station_names))


def convert_bytes(num):

    for x in ['bytes', 'KiB', 'MiB', 'GiB']:
        if num < 1024.0:
            return "%3.1f %s" % (num, x)
        num /= 1024.0


def format_elapsed_time(seconds):

    if seconds < 60:
        return f"{seconds:.3f} seconds"
    elif seconds < 3600:
        minutes, seconds = divmod(seconds, 60)
        return f"{int(minutes)} minutes {int(seconds)} seconds"
    else:
        hours, remainder = divmod(seconds, 3600)
        minutes, seconds = divmod(remainder, 60)
        if minutes == 0:
            return f"{int(hours)} hours {int(seconds)} seconds"
        else:
            return f"{int(hours)} hours {int(minutes)} minutes {int(seconds)} seconds"


def estimate_file_size(weather_station_names, num_rows_to_create):

    total_name_bytes = sum(len(s.encode("utf-8")) for s in weather_station_names)
    avg_name_bytes = total_name_bytes / float(len(weather_station_names))

    avg_temp_bytes = 4.400200100050025

    avg_line_length = avg_name_bytes + avg_temp_bytes + 2

    human_file_size = convert_bytes(num_rows_to_create * avg_line_length)

    return f"Estimated max file size is:  {human_file_size}."


def build_test_data(weather_station_names, num_rows_to_create):

    start_time = time.time()
    coldest_temp = -99.9
    hottest_temp = 99.9
    station_names_10k_max = random.choices(weather_station_names, k=10_000)
    batch_size = 10000 
    chunks = num_rows_to_create // batch_size
    print('building test data...')

    try:
        with open("../../../data/measurements.txt", 'w') as file:
            progress = 0
            for chunk in range(chunks):
                
                batch = random.choices(station_names_10k_max, k=batch_size)
                prepped_deviated_batch = '\n'.join([f"{station};{random.uniform(coldest_temp, hottest_temp):.1f}" for station in batch]) # :.1f should quicker than round on a large scale, because round utilizes mathematical operation
                file.write(prepped_deviated_batch + '\n')
                
                if (chunk + 1) * 100 // chunks != progress:
                    progress = (chunk + 1) * 100 // chunks
                    bars = '=' * (progress // 2)
                    sys.stdout.write(f"\r[{bars:<50}] {progress}%")
                    sys.stdout.flush()
        sys.stdout.write('\n')
    except Exception as e:
        print("error: ")
        print(e)
        exit()
    
    end_time = time.time()
    elapsed_time = end_time - start_time
    file_size = os.path.getsize("../../../data/measurements.txt")
    human_file_size = convert_bytes(file_size)
 
    print("success")
    print(f"Actual file size:  {human_file_size}")
    print(f"Elapsed time: {format_elapsed_time(elapsed_time)}")


def main():
    check_args(sys.argv)
    num_rows_to_create = int(sys.argv[1])
    weather_station_names = []
    weather_station_names = build_weather_station_name_list()
    print(estimate_file_size(weather_station_names, num_rows_to_create))
    build_test_data(weather_station_names, num_rows_to_create)
    print("complete")


if __name__ == "__main__":
    main()
exit()