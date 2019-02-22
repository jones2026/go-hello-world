FROM scratch
WORKDIR /app
COPY ./static ./static
COPY ./templates ./templates
COPY ./main ./
EXPOSE 80
CMD ["./main"]
