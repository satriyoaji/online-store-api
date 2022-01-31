# EVERMOS ONLINE STORE CASE #

### PROBLEM ###
We are members of the engineering team of an online store. When we look at ratings for our online store application, we received the following
facts:<br>
- Customers were able to put items in their cart, check out, and then pay. After several days, many of our customers received calls from<br>
our Customer Service department stating that their orders have been canceled due to stock unavailability.<br>
- These bad reviews generally come within a week after our 12.12 event, in which we held a large flash sale and set up other major
discounts to promote our store.<br>
<br>
After checking in with our Customer Service and Order Processing departments, we received the following additional facts:<br>
- Our inventory quantities are often misreported, and some items even go as far as having a negative inventory quantity.<br>
- The misreported items are those that performed very well on our 12.12 event.<br>
- Because of these misreported inventory quantities, the Order Processing department was unable to fulfill a lot of orders, and thus
requested help from our Customer Service department to call our customers and notify them that we have had to cancel their orders.<br>

### EXPLANATION PROBLEM ###
I think those bad reviews during the 12.12 event is happened because maybe a lot of customer orders were canceled because they got information that the stock was unavailable.
After checking by the customer service and order processing departments, it turned out that the items that performed very well at the 12.12 event were frequently misreported in inventory quantities. There are even some that have negative inventory counts. 
So the Order Processing department was unable to fulfill a lot of orders, and the Customer Service department notify their customers that they had to cancel their orders.

### SOLUTION ###
We need to sync the general stock with item that will have been at customer's checkout, 
so we can do reduce general stock before customer add to cart or notify customer that stock in unavailable when checkout that item.
So if the general stock of that item is 0, customer won't able to add that item to cart.
Then if we do 12.12 event sale, when customer going to checkout that item, we can prevent it from out of stock, because there's item stock checking in API.

### Technical Solution (proof of concept) ###
#### HOW TO RUN API ####
1. clone this repository<br>
2. get all required library using command `go mod tidy` on root path<br>
3. copy **.env.example** to file **.env** and config it with your database (using mysql)<br>
4. run the main.go<br>
5. its will auto migrate the database from models and there are seeders to fill the database<br>

#### HOW TO RUN UNIT TEST (Create new order API) ####
1. on the root directory, run in command line
    ```bash
    go test -v -run TestConnection
    ```
   <br>
2. Then open new terminal, on the root directory, run in command line
    ```bash
    go test -v -run TestCreateNewOrder
    ```
   <br>
3. So you'll see the response of API Create New Order<br>

#### or you can run throughly the `test.http` file ####