# Creating the app from empty image
FROM scratch
# Making the default directory current
WORKDIR /usr/local/bin
# Copying the resulting build file
COPY main ./
# Opening the default port
EXPOSE 8080
# Executing the build
CMD ["main"]
