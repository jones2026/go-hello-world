FROM scratch
WORKDIR /app
COPY ./main ./
EXPOSE 8080
CMD ["./main"]
