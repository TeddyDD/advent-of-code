prev == "" {
    prev = $0
}
prev < $0 {
    count += 1
}
{
    prev = $0
}
END {
    print count
}
