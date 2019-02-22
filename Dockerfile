FROM scratch
WORKDIR /app
COPY ./main ./
COPY ./static ./static
COPY ./templates ./templates
EXPOSE 80
CMD ["./main"]
